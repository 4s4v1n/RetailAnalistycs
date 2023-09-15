import $api from "../http";
import { AxiosResponse } from "axios";
import { GroupsResponse } from "../models/response/GroupsResponse";

const tableName = "groups"

export default class GroupsService {
    static async get(): Promise<GroupsResponse[]> {
        const response = $api.get<GroupsResponse[]>(`/view/${tableName}`)
        return (await response).data;
    }

    static async export(): Promise<AxiosResponse> {
        const response = $api.get(`/export/${tableName}`)
        return response
    }
}