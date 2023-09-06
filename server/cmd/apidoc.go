package main

import (
	"log"

	_ "github.com/hojin-kr/go-http-game-server/cmd/docs"
	EtcModel "github.com/hojin-kr/go-http-game-server/cmd/model/etc"
	ProfileModel "github.com/hojin-kr/go-http-game-server/cmd/model/profile"
	SocialModel "github.com/hojin-kr/go-http-game-server/cmd/model/social"
	UserModel "github.com/hojin-kr/go-http-game-server/cmd/model/user"
)

func apidoc() {
	log.Println(UserModel.User{})
	log.Println(ProfileModel.Profile{})
	log.Println(EtcModel.Etc{})
	log.Println(SocialModel.Social{})
}

// User APIs

// @Summary Get user by UUID 앱부팅시 UUID를 생성하고, UUID를 통해 사용자 정보를 조회한다
// @Description 앱부팅시 UUID를 생성하고, UUID를 통해 사용자 정보를 조회한다. (초기화시 매번 호출)
// @Tags User
// @ID get-user-by-uuid
// @Accept  json
// @Produce  json
// @Param uuid path string false "UUID"
// @Success 200 {object} UserModel.User
// @Router /api/v1/user/{uuid} [get]
func GetUserByUUID() {}

// @Summary Get recovery code by user_id 사용자의 복구 코드를 조회한다
// @Description 사용자의 복구 코드를 조회한다.
// @Tags User
// @ID get-recovery-code-by-user-id
// @Accept  json
// @Produce  json
// @Param user_id path int false "User ID"
// @Success 200 {object} UserModel.Recovery
// @Router /api/v1/user/{user_id}/recovery [get]
func GetRecoveryCodeByUserID() {}

// @Summary Recovery user 사용자의 복구 코드를 통해 UUID를 변경한다.
// @Description 사용자의 복구 코드를 통해 UUID를 변경한다.
// @Tags User
// @ID recovery-user
// @Accept  json
// @Produce  json
// @Param recovery body UserModel.RecoveryRequest true "Recovery"
// @Success 200 {object} UserModel.User
// @Router /api/v1/user/recovery [post]
func RecoveryUser() {}

// Profile APIs

// @Summary Get profile by ID 사용자의 프로필 정보를 조회한다
// @Description 사용자의 프로필 정보를 조회한다.
// @Tags Profile
// @ID get-profile-by-id
// @Accept  json
// @Produce  json
// @Param id path int false "ID"
// @Success 200 {object} ProfileModel.Profile
// @Router /api/v1/profile/{id} [get]
func GetProfileByID() {}

// @Summary Update profile nicname by ID 사용자의 프로필 닉네임 정보를 업데이트한다
// @Description 사용자의 프로필 닉네임 정보를 업데이트한다. 중복시 에러
// @Tags Profile
// @ID update-profile-nicname-by-id
// @Accept  json
// @Produce  json
// @Param profile body ProfileModel.Profile true "Profile"
// @Success 200 {object} ProfileModel.Profile
// @Router /api/v1/profile [post]
func UpdateProfileNicnameByID() {}

// Etc APIs

// @Summary Get etc by user_id and key 사용자의 기타 정보를 조회한다
// @Description 사용자의 기타 정보를 조회한다.
// @Tags Etc
// @ID get-etc-by-user-id-and-key
// @Accept  json
// @Produce  json
// @Param user_id path int false "User ID"
// @Param key path string false "Key"
// @Success 200 {object} EtcModel.Etc
// @Router /api/v1/etc/{user_id}/{key} [get]
func GetEtcByUserID() {}

// @Summary Get all etc by user_id 사용자의 모든 기타 정보를 조회한다
// @Description 사용자의 모든 기타 정보를 조회한다.
// @Tags Etc
// @ID get-all-etc-by-user-id
// @Accept  json
// @Produce  json
// @Param user_id path int false "User ID"
// @Success 200 {array} EtcModel.Etc
// @Router /api/v1/etc/{user_id} [get]
func GetAllEtcByUserID() {}

// @Summary Update etc by user_id 사용자의 기타 정보를 업데이트한다
// @Description 사용자의 기타 정보를 업데이트한다.
// @Tags Etc
// @ID update-etc-by-user-id
// @Accept  json
// @Produce  json
// @Param etc body EtcModel.Etc true "Etc"
// @Success 200 {object} EtcModel.Etc
// @Router /api/v1/etc [post]
func UpdateEtcByUserID() {}

// Social APIs

// @Summary Get social by target_id and type 타겟의 소셜 정보를 조회한다
// @Description 타겟의 소셜 정보를 조회한다.
// @Tags Social
// @ID get-social-by-target-id-and-type
// @Accept  json
// @Produce  json
// @Param target_id path int false "Target ID"
// @Param type path string false "Type"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} SocialModel.Social
// @Router /api/v1/social/{target_id}/{type}/{limit}/{offset} [get]
func GetSocialByTargetIDAndTypeLimitOffset() {}

// @Summary Get social count by target_id and type 타겟의 소셜 정보 카운트를 조회한다
// @Description 타겟의 소셜 정보 카운트를 조회한다.
// @Tags Social
// @ID get-social-count-by-target-id-and-type
// @Accept  json
// @Produce  json
// @Param target_id path int false "Target ID"
// @Param type path string false "Type"
// @Success 200 {int} int
// @Router /api/v1/social/{target_id}/{type} [get]
func GetSocialCountByTargetIDAndType() {}

// @Summary Get social by user_id and type 사용자의 소셜 정보를 조회한다
// @Description 사용자의 소셜 정보를 조회한다.
// @Tags Social
// @ID get-social-by-user-id-and-type
// @Accept  json
// @Produce  json
// @Param user_id path int false "User ID"
// @Param type path string false "Type"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} SocialModel.Social
// @Router /api/v1/social/{user_id}/{type}/{limit}/{offset} [get]
func GetSocialByUserIDAndTypeLimitOffset() {}

// @Summary Insert social 사용자의 소셜 정보를 추가한다
// @Description 사용자의 소셜 정보를 추가한다.
// @Tags Social
// @ID insert-social
// @Accept  json
// @Produce  json
// @Param social body SocialModel.SocialRequest true "Social"
// @Success 200 {object} SocialModel.Social
// @Router /api/v1/social [post]
func InsertSocial() {}

// @Summary Delete social by ID 소셜 정보를 삭제한다
// @Description 소셜 정보를 삭제한다.
// @Tags Social
// @ID delete-social-by-id
// @Accept  json
// @Produce  json
// @Param social body SocialModel.SocialDeleteRequest true "Social"
// @Success 200 {object} SocialModel.Social
// @Router /api/v1/social/delete [post]
func DeleteSocialByID() {}
