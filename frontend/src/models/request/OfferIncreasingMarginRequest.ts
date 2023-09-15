export interface OfferIncreasingMarginRequest {
    count_group: number;
    max_churn_rate: number;
    max_stability_index: number;
    max_index_sku: number;
    margin_share: number;
};

export const DefaultOfferIncreasingMarginRequest: OfferIncreasingMarginRequest = {
    count_group: 10,
    max_churn_rate: 10,
    max_stability_index: 10,
    max_index_sku: 0.5,
    margin_share: 10,
};