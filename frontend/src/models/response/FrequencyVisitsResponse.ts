export interface FrequencyVisitsResponse {
    customer_id: number;
    start_date: string;
    end_date: string;
    required_transactions_count: number;
    group_name: string;
    offer_discount_depth: number;
}