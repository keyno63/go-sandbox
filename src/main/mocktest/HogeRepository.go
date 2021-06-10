//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package mocktest

// Factory : Creates objects in this package
type HogeRepository interface {
	Get() (User, error)
}

type HogeRepositoryImpl struct {
}

func (r *HogeRepositoryImpl) Get(userId string) (User, error) {
	return User{
		ID: "someId",
	}, nil
}
