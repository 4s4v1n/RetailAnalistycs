import { FC, useState } from "react";
import { observer } from "mobx-react-lite";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { notifyError, notifySucces } from "../components/Notify/Notify";
import { saveAs } from 'file-saver'
import GrowthAverageService from "../services/GrowthAverageService";
import HeaderForm from "../components/FunctionHeader/FunctionHeader";
import { FUNCTION_GROWTH } from "../utils/const";
import View from "../components/View/View";
import { GrowthAverageResponse } from "../models/response/GrowthAverageResponse";

const GrowthAverage: FC = () => {

    const [data, setData] = useState<GrowthAverageResponse[]>();

    const queryClient = useQueryClient()

    const name = FUNCTION_GROWTH

    const { mutate: mutationForm } = useMutation({
        mutationFn: GrowthAverageService.get,
        onSuccess: (data) => {
            setData(data)
            notifySucces(`Successful response`)
            queryClient.invalidateQueries({ queryKey: [{ name }] })
        },
        onError: () => {
            notifyError(`Can't response value in ${name}`);
        },
    })

    const { mutate: mutateExport } = useMutation({
        mutationFn: GrowthAverageService.export,
        onSuccess: (data) => {
            notifySucces(`Successful export`)
            const blob = new Blob([data.data], { type: 'text/csv;charset=utf-8;' });
            saveAs(blob, `${name}.csv`);
        },
        onError: () => {
            notifyError(`Can't export ${name}`);
        },
    })


    return (
        <div>
            <HeaderForm mutationForm={mutationForm} mutationExport={mutateExport} typeFunction={FUNCTION_GROWTH} />
            <View data={data} />
        </div>
    )
}

export default observer(GrowthAverage);