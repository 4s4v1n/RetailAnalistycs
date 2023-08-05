import { FC, useState } from "react";
import { observer } from "mobx-react-lite";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { notifyError, notifySucces } from "../components/Notify/Notify";
import { saveAs } from 'file-saver'
import HeaderForm from "../components/FunctionHeader/FunctionHeader";
import { FUNCTION_FREQUENCY } from "../utils/const";
import View from "../components/View/View";
import { FrequencyVisitsResponse } from "../models/response/FrequencyVisitsResponse";
import FrequencyVisitsService from "../services/FrequencyVisitsService";

const FreuencyVisits: FC = () => {

    const [data, setData] = useState<FrequencyVisitsResponse[]>();

    const queryClient = useQueryClient()

    const name = FUNCTION_FREQUENCY

    const { mutate: mutationForm } = useMutation({
        mutationFn: FrequencyVisitsService.get,
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
        mutationFn: FrequencyVisitsService.export,
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
            <HeaderForm mutationForm={mutationForm} mutationExport={mutateExport} typeFunction={FUNCTION_FREQUENCY} />
            <View data={data} />
        </div>
    )
}

export default observer(FreuencyVisits);