export interface PurchaiseHistoryResponse {
    customer_id: number;
    transaction_id: number;
    transaction_datetime: string;
    group_id: number;
    group_cost: number;
    group_sum: number;
    group_sum_paid: number;
}