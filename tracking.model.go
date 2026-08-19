package types

import "time"

// TrackingSource adalah asal data AIS/posisi kapal.
// Sumber dapat berasal dari internal PLN (IP), sistem BAG, provider eksternal
// (MarineTraffic), atau input manual operator.
type TrackingSource string

const (
	TrackingSourceIP            TrackingSource = "IP"
	TrackingSourceBAG           TrackingSource = "BAG"
	TrackingSourceMarineTraffic TrackingSource = "MarineTraffic"
	TrackingSourceManual        TrackingSource = "Manual"
)

// ShipType klasifikasi kapal untuk filter UI Tracking Kapal.
// Nilai mengikuti kategori yang sudah dipakai frontend (Vessel/TUG Boat/SPB).
type ShipType string

const (
	ShipTypeVessel  ShipType = "Vessel"
	ShipTypeTugBoat ShipType = "TUG Boat"
	ShipTypeSPB     ShipType = "SPB"
)

// VesselTracking berisi data AIS/navigasi kapal.
//
// Sumber: MarineTraffic / IP / BAG / provider AIS lainnya.
//
// Catatan tipe data:
//   - Field numerik (Latitude, Longitude, Speed, Course, Heading) memakai
//     pointer agar null (data AIS tidak tersedia) dapat dibedakan dari 0 yang
//     merupakan nilai valid (mis. Speed=0 = kapal berhenti, AISStatus=0 =
//     "under way using engine"). Jangan mengisi field kosong dengan 0.
//   - MMSI & IMO tetap string agar leading-zero aman dan dapat menampung
//     sentinel "-" dari sumber lama (akan dinormalisasi ke null oleh adapter).
//   - AISStatus adalah navigational status AIS (kode 0-15 per spec AIS), BUKAN
//     status bisnis pengiriman. Status bisnis ada di ShipmentTracking.Status.
type VesselTracking struct {
	MMSI      string         `json:"mmsi"`
	IMO       string         `json:"imo"`
	ShipID    *string        `json:"ship_id"`    // ID internal MarineTraffic; null dari sumber lain
	Name      string         `json:"name"`       // nama kapal bersih (tanpa suffix tug)
	ShipType  ShipType       `json:"ship_type"`  // Vessel/TUG Boat/SPB
	Latitude  *float64       `json:"latitude"`   // derajat desimal, range -90..90
	Longitude *float64       `json:"longitude"`  // derajat desimal, range -180..180
	Speed     *float64       `json:"speed"`      // knot, null = tidak diketahui
	Course    *float64       `json:"course"`     // derajat 0-359.9
	Heading   *float64       `json:"heading"`    // derajat 0-359
	AISStatus *int           `json:"ais_status"` // kode AIS 0-15, BUKAN status bisnis
	Timestamp *time.Time     `json:"timestamp"`  // ISO 8601 dengan offset
	Source    TrackingSource `json:"source"`
}

// ShipmentTracking berisi data bisnis/logistik pengiriman.
//
// Status di sini adalah status proses pengiriman aplikasi (BERANGKAT/TIBA/
// BONGKAR/dll), BUKAN navigational status AIS. Jangan disamakan dengan
// VesselTracking.AISStatus.
//
// Field koordinat port/destination disediakan agar frontend dapat menggambar
// garis asal-tujuan di peta tanpa lookup tambahan.
type ShipmentTracking struct {
	IDPengiriman           string     `json:"id_pengiriman"`
	NoJadwal               string     `json:"no_jadwal"`
	VesselName             string     `json:"vessel_name"` // string mentah dari sumber lama ("MV. X / -")
	VolumeBL               *float64   `json:"volume_bl"`   // ton/Mt, null jika belum BL
	ETA                    *time.Time `json:"eta"`         // date only (YYYY-MM-DD)
	PortOfLoading          string     `json:"port_of_loading"`
	PortOfLoadingLatitude  *float64   `json:"port_of_loading_latitude"`
	PortOfLoadingLongitude *float64   `json:"port_of_loading_longitude"`
	Destination            string     `json:"destination"`
	DestinationLatitude    *float64   `json:"destination_latitude"`
	DestinationLongitude   *float64   `json:"destination_longitude"`
	Status                 string     `json:"status"` // enum bisnis: BERANGKAT/TIBA/BONGKAR/...
	Flag                   string     `json:"flag"`
	LastUpdate             *time.Time `json:"last_update"` // ISO 8601
}

// TrackingItem adalah satu entri tracking = vessel + shipment.
// Pemisahan vessel vs shipment memungkinkan merge data AIS real-time dari
// provider eksternal (MarineTraffic) dengan data bisnis dari sistem internal
// (IP/BAG) secara independen.
type TrackingItem struct {
	Vessel   VesselTracking   `json:"vessel"`
	Shipment ShipmentTracking `json:"shipment"`
}

// TrackingTypeSummary agregat per ship_type untuk legend UI Tracking Kapal.
type TrackingTypeSummary struct {
	Count       int     `json:"count"`
	TotalVolume float64 `json:"total_volume_bl"`
}

// TrackingListMeta metadata agregat untuk halaman list Tracking Kapal.
// Frontend dapat langsung memakai ByType untuk merender legend
// (Total/Vessel/TUG Boat/SPB) tanpa menghitung ulang client-side.
//
// DataStale=true berarti data berasal dari snapshot terakhir di
// tracking_history (upstream sedang tidak bisa di-fetch). LastUpdatedAt
// adalah waktu snapshot terbaru yang dipakai (ISO 8601); null jika data
// live (upstream sehat).
type TrackingListMeta struct {
	Total         int                              `json:"total"`
	TotalVolume   float64                          `json:"total_volume_bl"`
	ByType        map[ShipType]TrackingTypeSummary `json:"by_type"`
	DataStale     bool                             `json:"data_stale"`
	LastUpdatedAt *time.Time                       `json:"last_updated_at"`
}

// TrackingListResponse response untuk endpoint list tracking kapal.
type TrackingListResponse struct {
	Data []TrackingItem   `json:"data"`
	Meta TrackingListMeta `json:"meta"`
}

// TrackingDetailResponse response untuk endpoint detail satu pengiriman.
type TrackingDetailResponse struct {
	Data TrackingItem `json:"data"`
}

// TrackingHistory adalah snapshot posisi + status kapal pada satu titik waktu.
// Tabel ini diisi oleh polling job yang fetch dari sumber legacy (IP) tiap
// TRACKING_POLL_INTERVAL menit. Setiap row = satu checkpoint per voyage.
//
// Voyage diidentifikasi oleh IDPengiriman. Satu kapal (IMO/MMSI) dapat punya
// banyak voyage berbeda dari waktu ke waktu — masing-masing punya IDPengiriman,
// NoJadwal, dan Destination yang berbeda. Voyage yang sudah bongkar (status
// masuk daftar TRACKING_CLOSED_STATUS) tidak muncul di list realtime lagi,
// tapi history-nya tetap tersimpan di tabel ini untuk analytics.
type TrackingHistory struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	IDPengiriman  string     `gorm:"type:varchar(64);index:idx_tracking_history_pengiriman_time,priority:1;not null" json:"id_pengiriman"`
	NoJadwal      string     `gorm:"type:varchar(64)" json:"no_jadwal"`
	VesselName    string     `gorm:"type:varchar(128)" json:"vessel_name"`
	IMO           string     `gorm:"type:varchar(32)" json:"imo"`
	MMSI          string     `gorm:"type:varchar(32)" json:"mmsi"`
	Latitude      *float64   `gorm:"type:double precision" json:"latitude"`
	Longitude     *float64   `gorm:"type:double precision" json:"longitude"`
	Speed         *float64   `gorm:"type:double precision" json:"speed"`
	Course        *float64   `gorm:"type:double precision" json:"course"`
	VolumeBL      *float64   `gorm:"type:double precision" json:"volume_bl"`
	ETA           *time.Time `gorm:"type:timestamptz" json:"eta"`
	PelabuhanMuat string     `gorm:"type:varchar(128)" json:"pelabuhan_muat"`
	Tujuan        string     `gorm:"type:varchar(128)" json:"tujuan"`
	StatusKapal   string     `gorm:"type:varchar(64);index:idx_tracking_history_status" json:"status_kapal"`
	Flag          string     `gorm:"type:varchar(32)" json:"flag"`
	Source        string     `gorm:"type:varchar(16)" json:"source"`
	CapturedAt    time.Time  `gorm:"type:timestamptz;index:idx_tracking_history_pengiriman_time,priority:2;not null;default:CURRENT_TIMESTAMP" json:"captured_at"`
}

// TableName override nama tabel GORM.
func (TrackingHistory) TableName() string { return "tracking_history" }

// TrackingHistoryResponse response untuk endpoint history satu voyage.
type TrackingHistoryResponse struct {
	IDPengiriman string            `json:"id_pengiriman"`
	VesselName   string            `json:"vessel_name"`
	NoJadwal     string            `json:"no_jadwal"`
	Tujuan       string            `json:"tujuan"`
	History      []TrackingHistory `json:"history"`
}

// TrackingHistoryListResponse response untuk endpoint list semua history
// (dengan filter opsional: date range, vessel, destination).
type TrackingHistoryListResponse struct {
	Data []TrackingHistorySummary `json:"data"`
}

// TrackingHistorySummary ringkasan satu voyage untuk list history.
type TrackingHistorySummary struct {
	IDPengiriman  string     `json:"id_pengiriman"`
	NoJadwal      string     `json:"no_jadwal"`
	VesselName    string     `json:"vessel_name"`
	IMO           string     `json:"imo"`
	MMSI          string     `json:"mmsi"`
	PelabuhanMuat string     `json:"pelabuhan_muat"`
	Tujuan        string     `json:"tujuan"`
	VolumeBL      *float64   `json:"volume_bl"`
	StatusKapal   string     `json:"status_kapal"`
	FirstSeen     *time.Time `json:"first_seen"`
	LastSeen      *time.Time `json:"last_seen"`
	SnapshotCount int        `json:"snapshot_count"`
}
