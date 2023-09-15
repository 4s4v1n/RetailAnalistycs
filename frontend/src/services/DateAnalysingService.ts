import $api from "../http";
import { AxiosResponse } from "axios";
import { DateAnalysingResponse } from "../models/response/DateAnalysingResponse";
import { DateOfAnalysingRequest } from "../models/request/DateOfAnalysingRequest";

const tableName = "date_of_analysing_formation"

export default class DateAnalysingService {
    static async get(): Promise<DateAnalysingResponse[]> {
        const response = $api.get<DateAnalysingResponse[]>(`/table/${tableName}`)
        return (await response).data;
    }

    static async patch(data: DateOfAnalysingRequest): Promise<AxiosResponse> {
        console.log(data)
        const response = $api.patch(`/table/${tableName}`, data)
        return response
    }
}
