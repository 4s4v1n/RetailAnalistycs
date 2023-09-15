export interface PeriodsResponse {
    customer_id: number;
    group_id: number;
    first_group_purchase_date: string;
    last_group_purchase_date: string;
    group_purchase: string;
    group_frequency: number;
    group_min_discount: number;
}