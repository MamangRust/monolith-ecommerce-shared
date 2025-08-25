package recordmapper

import (
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/record"
)

type merchantSocialLinkRecordMapper struct {
}

func NewMerchantSocialLinkRecordMapper() *merchantSocialLinkRecordMapper {
	return &merchantSocialLinkRecordMapper{}
}

func (m *merchantSocialLinkRecordMapper) ToMerchantSocialLinkRecord(link *db.MerchantSocialMediaLink) *record.MerchantSocialLinkRecord {

	return &record.MerchantSocialLinkRecord{
		ID:               int(link.MerchantSocialID),
		MerchantDetailID: int(link.MerchantDetailID),
		Platform:         link.Platform,
		URL:              link.Url,
		CreatedAt:        link.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:        link.UpdatedAt.Time.Format("2006-01-02"),
	}
}

func (m *merchantSocialLinkRecordMapper) ToMerchantSocialLinksRecord(links []*db.MerchantSocialMediaLink) []*record.MerchantSocialLinkRecord {
	var records []*record.MerchantSocialLinkRecord
	for _, link := range links {
		records = append(records, m.ToMerchantSocialLinkRecord(link))
	}
	return records
}
