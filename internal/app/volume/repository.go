package volume

type Repository interface {
	Create(volume Volume) Volume
	List() []Volume
	Get(id string) (Volume, error)
	Update(volume Volume) (Volume, error)
	Delete(id string) error
	Reset()
}

type SnapshotRepository interface {
	Create(snapshot Snapshot) Snapshot
	List() []Snapshot
	Get(id string) (Snapshot, error)
	Delete(id string) error
	Reset()
}

type TransferRepository interface {
	Create(transfer Transfer) Transfer
	List() []Transfer
	Get(id string) (Transfer, error)
	Delete(id string) error
	Reset()
}
