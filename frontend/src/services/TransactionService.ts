import $api from "../http";
import { AxiosResponse } from "axios";
import { TransactionResponse } from "../models/response/TransactionResponse";

const tableName = "transactions"

export default class TransactionService {
    static async get(): Promise<TransactionResponse[]> {
        const response = $api.get<TransactionResponse[]>(`/table/${tableName}`)
        return (await response).data;
    }

    static async post(data: TransactionResponse): Promise<AxiosResponse> {
        data.customer_card_id = Number(data.customer_card_id);
        data.transaction_id = Number(data.transaction_id);
        data.transaction_sum = Number(data.transaction_sum);
        data.transaction_store_id = Number(data.transaction_store_id);
        const response = $api.post(`/table/${tableName}`, data)
        return response
    }

    static async patch(data: TransactionResponse): Promise<AxiosResponse> {
        data.customer_card_id = Number(data.customer_card_id);
        data.transaction_id = Number(data.transaction_id);
        data.transaction_sum = Number(data.transaction_sum);
        data.transaction_store_id = Number(data.transaction_store_id);
        const response = $api.patch(`/table/${tableName}`, data)
        return response
    }

    static async delete(data: string): Promise<AxiosResponse> {
        const response = $api.delete(`/table/${tableName}/${data}`)
        return response
    }

    static async export(): Promise<AxiosResponse> {
        const response = $api.get(`/export/${tableName}`)
        return response
    }

    static async import(data: any): Promise<AxiosResponse> {
        for (let i = 0; i < data.length; i++) {
            data[i].customer_card_id = Number(data[i].customer_card_id);
            data[i].transaction_id = Number(data[i].transaction_id);
            data[i].transaction_sum = Number(data[i].transaction_sum);
            data[i].transaction_store_id = Number(data[i].transaction_store_id);
        }
        const response = $api.post(`/import/${tableName}`, data)
        return response
    }
}
