export interface GrowthAverageRequest {
    method: string;
    first: string;
    last: string;
    number: number;
    coefficient: number;
    max_churn_rate: number;
    max_discount_share: number;
    margin_share: number;
}

export const DefaultGrowthAverageRequest: GrowthAverageRequest = {
    method: 'PERIOD',
    first: '01.01.2018',
    last: '12.12.2022',
    number: 1000,
    coefficient: 1.45,
    max_churn_rate: 12,
    max_discount_share: 1,
    margin_share: 100,
};

