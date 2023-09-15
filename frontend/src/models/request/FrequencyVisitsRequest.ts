export interface FrequencyVisitsRequest {
    first: string;
    last: string;
    value_transaction: number;
    max_churn_rate: number;
    max_discount_share: number;
    margin_share: number;
}

export const DefaultFrequencyVisitsRequest: FrequencyVisitsRequest = {
    first: '01.01.2018',
    last: '12.12.2022',
    value_transaction: 30,
    max_churn_rate: 12,
    max_discount_share: 1,
    margin_share: 100,
};