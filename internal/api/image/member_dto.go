package image

import appimage "github.com/JSYoo5B/SandStack/internal/app/image"

type createMemberRequest struct {
	Member string `json:"member"`
}

type updateMemberRequest struct {
	Status string `json:"status"`
}

type memberListResponse struct {
	Members []memberDocument `json:"members"`
}

type memberDocument struct {
	ImageID   string `json:"image_id"`
	MemberID  string `json:"member_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Schema    string `json:"schema"`
}

func toMemberDocuments(members []appimage.Member) []memberDocument {
	documents := make([]memberDocument, 0, len(members))
	for _, member := range members {
		documents = append(documents, toMemberDocument(member))
	}

	return documents
}

func toMemberDocument(member appimage.Member) memberDocument {
	return memberDocument{
		ImageID:   member.ImageID,
		MemberID:  member.MemberID,
		Status:    member.Status,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
		Schema:    "/v2/schemas/member",
	}
}
