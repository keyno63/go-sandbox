package mocktest

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_hogeServiceImpl_Fuga(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 生成
	mocker := NewMockHogeRepository(ctrl)
	user := User{"mockedId"}
	mocker.EXPECT().Get().Return(user, nil)

	//mocker := mock.NewMockUserRepository(mocka)
	type fields struct {
		UserRepository *MockHogeRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    *User
		wantErr bool
	}{
		{
			name: "hoge",
			fields: fields{
				mocker,
			},
			want: &User{
				"mockedId",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &hogeServiceImpl{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := s.Fuga()
			if (err != nil) != tt.wantErr {
				t.Errorf("Fuga() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fuga() got = %v, want %v", got, tt.want)
			}
		})
	}
}
