package repository

import (
	requestmodel "Laptop_Lounge/pkg/models/requestModel"
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Test GetUserDetails
func TestGetUserDetails(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    *requestmodel.UserDetails
		wantErr error
	}{
		{
			name: "succesfully got result",
			args: "1",
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, name , email, phone, referal_code FROM users WHERE id= ?").
					WillReturnRows(s.NewRows([]string{"id", "name", "email", "phone", "referal_code"}).
						AddRow("1", "Shahabaz", "shahabaz@gmail.com", "9496703880", "aabb2"))
			},
			want: &requestmodel.UserDetails{
				Id:          "1",
				Name:        "Shahabaz",
				Email:       "shahabaz@gmail.com",
				Phone:       "9496703880",
				ReferalCode: "aabb2",
			},
			wantErr: nil,
		}, {
			name: "No User Found In This ID",
			args: "1",
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, name , email, phone, referal_code FROM users WHERE id= ?").
					WillReturnRows(s.NewRows([]string{"id", "name", "email", "phone", "referal_code"}))
			},
			want:    nil,
			wantErr: errors.New("sorry, we couldn't find any data that matches your criteria in the database."),
		},
	}

	for _, tt := range tests {

		// Setup mock DB and gorm DB
		mockDB, mockSql, _ := sqlmock.New()
		defer mockDB.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})

		// Apply the stub
		tt.stub(mockSql)

		// Create the repository
		userRepository := NewUserRepository(DB)

		// Execute the method
		result, _ := userRepository.GetProfile(tt.args)

		assert.Equal(t, tt.want, result)
		//assert.Equal(t, tt.wantErr, result)

		// assert.Equal(t, tt.want, result)
		// 	if tt.wantErr != nil {
		// 		assert.Error(t, err)
		// 		assert.EqualError(t, err, tt.wantErr.Error())
		// 	} else {
		// 		assert.NoError(t, err)
		// 	}
	}
}

// Test Fetch Password using Phone number
func TestFetchPasswordUsingPhone(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    string
		wantErr error
	}{{
		name: "succefully Fetch Password",
		args: "9496703880",
		stub: func(s sqlmock.Sqlmock) {
			s.ExpectQuery(regexp.QuoteMeta("SELECT password FROM users WHERE phone=$1 AND status='active'")).
				WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("asdsfdfbdght345dgf"))
		},
		want:    "asdsfdfbdght345dgf",
		wantErr: nil,
	}, {
		name: "no user is exist",
		args: "9496703880",
		stub: func(s sqlmock.Sqlmock) {
			s.ExpectQuery(regexp.QuoteMeta("SELECT password FROM users WHERE phone=$1 AND status='active'")).
				WillReturnError(errors.New("no user exist or you are block by admin"))
		},
		want:    "",
		wantErr: errors.New("no user exist or you are blocked by admin"),
	},
	}

	for _, tt := range tests {
		mockDB, mockSql, _ := sqlmock.New()
		defer mockDB.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})

		tt.stub(mockSql)

		userRepository := NewUserRepository(DB)
		result, _ := userRepository.FetchPasswordUsingPhone(tt.args)

		assert.Equal(t, tt.want, result)

		//assert.Equal(t,tt.wantErr,result)

	}
}

// Test User Exist or Not
func TestUserExist(t *testing.T) {
	mockDB, mockSQL, _ := sqlmock.New()

	DB, _ := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})

	args := "9496703880"
	want := 1

	expectedQuery := "SELECT COUNT(*) FROM users WHERE phone=$1 AND status!=$2"
	mockSQL.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WillReturnRows(sqlmock.NewRows([]string{"phone"}).AddRow("1"))

	userRepository := NewUserRepository(DB)
	result := userRepository.IsUserExist(args)

	assert.Equal(t, want, result)
}

// Test Creating User
func TestUserSignUp(t *testing.T) {
	tests := []struct {
		name    string
		args    *requestmodel.UserDetails
		stub    func(sqlmock.Sqlmock)
		want    *responsemodel.SignupData
		wantErr error
	}{
		{
			name: "Successfully User Signup",
			args: &requestmodel.UserDetails{
				Name:        "Shahabaz",
				Email:       "shahabaz@gmail.com",
				Phone:       "9496703880",
				ReferalCode: "aabb2",
				Password:    "8Ubxsdah786dbsjkhdshlsdhc668ms",
			},
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name,email,phone,password,referal_code) values($1,$2,$3,$4,$5) RETURNING *")).
					WillReturnRows(s.NewRows([]string{"id", "name", "email", "phone", "referal_code"}).
						AddRow("1","Shahabaz", "shahabaz@gmail.com", "9496703880", "aabb2"))
			},
			want: &responsemodel.SignupData{
				ID:          "1",
				Name:        "Shahabaz",
				Email:       "shahabaz@gmail.com",
				Phone:       "9496703880",
				ReferalCode: "aabb2",
			},
			wantErr: nil,
		},
		{
			name: "Error at user creation",
			args: &requestmodel.UserDetails{
				Name:        "Shahabaz",
				Email:       "shahabaz@gmail.com",
				Phone:       "9496703880",
				ReferalCode: "aabb2",
				Password:    "8Ubxsdah786dbsjkhdshlsdhc668ms",
			},
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name,email,phone,password,referal_code) values($1,$2,$3,$4,$5) RETURNING *")).
					WillReturnError(errors.New("data missmatching can't store in database"))
			},
			want:    nil,
			wantErr: errors.New("data missmatching can't store in database"),
		},
	}

	for _, tt := range tests {
		mockDB, mockSql, _ := sqlmock.New()
		defer mockDB.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})

		tt.stub(mockSql)

		userRepository := NewUserRepository(DB)
		result, _ := userRepository.UserSignUp(tt.args)

		assert.Equal(t, tt.want, result)
		// Uncomment the next line if you want to compare errors as well
		// assert.Equal(t, tt.wantErr, err)
	}
}

// test For Creating Address
func TestCreateAddress(t *testing.T) {
	test := []struct {
		name    string
		args    *requestmodel.Address
		stub    func(sqlmock.Sqlmock)
		want    *requestmodel.Address
		wantErr error
	}{
		{
			name: "Create user address",
			args: &requestmodel.Address{
				Userid:      "123",
				FirstName:   "Emmanuel",
				LastName:    "Johnson",
				Street:      "Edapally",
				City:        "Ponnekara",
				State:       "kerala",
				Pincode:     "567843",
				LandMark:    "Near Park",
				PhoneNumber: "7866778892",
			},
			stub: func(sqlmock sqlmock.Sqlmock) {
				sqlmock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO addresses ( userid, first_name, last_name, street, city, state, pincode, land_mark, phone_number) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "userid", "first_name", "last_name", "street", "city", "state", "pincode", "land_mark", "phone_numbe"}).
						AddRow("1", "123", "Emmanuel", "Johnson", "Edapally", "Ponnekara", "kerala", "567843", "Near Park", "7866778892"))
			},
			want: &requestmodel.Address{
				ID:          "1",
				Userid:      "123",
				FirstName:   "Emmanuel",
				LastName:    "Johnson",
				Street:      "Edapally",
				City:        "Ponnekara",
				State:       "kerala",
				Pincode:     "567843",
				LandMark:    "Near Park",
				PhoneNumber: "7866778892",
			},
			wantErr: nil,
		}, {
			name: "error at addres creation of user",
			args: &requestmodel.Address{
				Userid:      "123",
				FirstName:   "Emmanuel",
				LastName:    "Johnson",
				Street:      "Edapally",
				City:        "Ponnekara",
				State:       "kerala",
				Pincode:     "567843",
				LandMark:    "Near Park",
				PhoneNumber: "7866778892",
			},
			stub: func(sqlmock sqlmock.Sqlmock) {
				sqlmock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO addresses ( userid, first_name, last_name, street, city, state, pincode, land_mark, phone_number) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "userid", "first_name", "last_name", "street", "city", "state", "pincode", "land_mark", "phone_numbe"}))
			},
			want:    nil,
			wantErr: errors.New("no data matching the specified criteria was found in the database"),
		},
	}

	for _, tt := range test {
		mockDB, mocksql, _ := sqlmock.New()
		defer mockDB.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})

		userRepository := NewUserRepository(DB)

		tt.stub(mocksql)

		result, err := userRepository.CreateAddress(tt.args)

		assert.Equal(t, tt.want, result)
		assert.Equal(t, tt.wantErr, err)
	}
}
