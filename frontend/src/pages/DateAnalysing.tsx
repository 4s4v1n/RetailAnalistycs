import { FC } from "react";
import { observer } from "mobx-react-lite";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { notifyError, notifySucces } from "../components/Notify/Notify";
import DateAnalysingService from "../services/DateAnalysingService";
import TableOnce from "../components/TableOnce/TableOnce";

const DateAnalysing: FC = () => {

    const queryClient = useQueryClient()

    const name = "date_of_analysing_formation"

    const { data } = useQuery([{ name }], () => DateAnalysingService.get(),)


    const { mutate: mutationUpdate } = useMutation({
        mutationFn: DateAnalysingService.patch,
        onSuccess: () => {
            notifySucces(`Successful update`)
            queryClient.invalidateQueries({ queryKey: [{ name }] })
        },
        onError: () => {
            notifyError(`Can't update value in ${name}`);
        },
    })

    return (
        <div>
            <div className="flex justify-between sm:ml-64 p-4 pt-20">
                <h1 className="text-4xl font-extrabold dark:text-white">{name}</h1>
            </div>
            <TableOnce data={data} mutationUpdate={mutationUpdate}/>
        </div>
    )
}

export default observer(DateAnalysing);