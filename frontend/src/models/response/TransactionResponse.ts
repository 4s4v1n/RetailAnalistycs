export interface TransactionResponse {
    transaction_id: number;
    customer_card_id: number;
    transaction_sum: number;
    transaction_datetime: string;
    transaction_store_id: number;
}