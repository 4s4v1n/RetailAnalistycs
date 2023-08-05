package entity

type Customer struct {
	CustomerId                  int64   `json:"customer_id" csv:"customer_id" db:"customer_id"`
	CustomerAverageCheck        float64 `json:"customer_average_check" csv:"customer_average_check" db:"customer_average_check"`
	CustomerAverageCheckSegment string  `json:"customer_average_check_segment" csv:"customer_average_check_segment" db:"customer_average_check_segment"`
	CustomerFrequency           float64 `json:"customer_frequency" csv:"customer_frequency" db:"customer_frequency"`
	CustomerFrequencySegment    string  `json:"customer_frequency_segment" csv:"customer_frequency_segment" db:"customer_frequency_segment"`
	CustomerInactivePeriod      float64 `json:"customer_inactive_period" csv:"customer_inactive_period" db:"customer_inactive_period"`
	CustomerChurnRate           float64 `json:"customer_churn_rate" csv:"customer_churn_rate" db:"customer_churn_rate"`
	CustomerChurnSegment        string  `json:"customer_churn_segment" csv:"customer_churn_segment" db:"customer_churn_segment"`
	CustomerSegment             int32   `json:"customer_segment" csv:"customer_segment" db:"customer_segment"`
	CustomerPrimaryStore        int64   `json:"customer_primary_store" csv:"customer_primary_store" db:"customer_primary_store"`
}
