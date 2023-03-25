package repository

type MicroRepository interface {
	User() UserRepository
	Auth() AuthRepository
	Category() CategoryRepository
	Book() BookRepository
	Wishlist() WishlistRepository
	Rating() RatingRepository
}

type MicroRepositorysImplementation struct {
	UserRepository
	AuthRepository
	CategoryRepository
	BookRepository
	WishlistRepository
	RatingRepository
}

func NewRegisterMicroServiceImplementation(userRepository UserRepository, authRepository AuthRepository, categoryRepository CategoryRepository, bookRepository BookRepository, wishlistRepository WishlistRepository, ratingRepository RatingRepository) MicroRepository {
	return &MicroRepositorysImplementation{UserRepository: userRepository, AuthRepository: authRepository, CategoryRepository: categoryRepository, BookRepository: bookRepository, WishlistRepository: wishlistRepository, RatingRepository: ratingRepository}
}

func (m *MicroRepositorysImplementation) User() UserRepository {
	return m.UserRepository
}

func (m *MicroRepositorysImplementation) Auth() AuthRepository {
	return m.AuthRepository
}

func (m *MicroRepositorysImplementation) Category() CategoryRepository {
	return m.CategoryRepository
}

func (m *MicroRepositorysImplementation) Book() BookRepository {
	return m.BookRepository
}

func (m *MicroRepositorysImplementation) Wishlist() WishlistRepository {
	return m.WishlistRepository
}

func (m *MicroRepositorysImplementation) Rating() RatingRepository {
	return m.RatingRepository
}
