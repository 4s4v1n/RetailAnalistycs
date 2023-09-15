import { FC, useState } from 'react';
import FunctionModal from '../Modal/FunctionModal';
import { FUNCTION_FREQUENCY, FUNCTION_GROWTH } from '../../utils/const';
import { DefaultGrowthAverageRequest } from '../../models/request/GrowthAverageRequest';
import { DefaultFrequencyVisitsRequest } from '../../models/request/FrequencyVisitsRequest';
import { DefaultOfferIncreasingMarginRequest } from '../../models/request/OfferIncreasingMarginRequest';

export interface HeadersProps {
    mutationForm: any;
    mutationExport: any;
    typeFunction: string;
}

const HeaderForm: FC<HeadersProps> = ({mutationForm, mutationExport, typeFunction }) => {
    const [openModalForm, setOpenModalForm] = useState<string | undefined>("default");
    const [dataForm, setDataForm] = useState<{ [key: string]: string }>({});


    const getKeys = (typeFunction: string): any => {
        if (typeFunction === FUNCTION_GROWTH) {
            const data = Object.entries(DefaultGrowthAverageRequest);
            return data;
        } else if (typeFunction === FUNCTION_FREQUENCY) {
            const data = Object.entries(DefaultFrequencyVisitsRequest);
            return data;
        } else {
            const data = Object.entries(DefaultOfferIncreasingMarginRequest);
            return data;
        }
    }

    const exportForm = () => {
        mutationExport(dataForm)
    }

    return (
        <div className="flex justify-between sm:ml-64 p-4 pt-20">
            <FunctionModal openModalForm={openModalForm} setOpenModalForm={setOpenModalForm} keys={getKeys(typeFunction)} mutationForm={mutationForm} setDataForm={setDataForm} />
            <h1 className="text-4xl font-extrabold dark:text-white">{typeFunction}</h1>
            <div className="flex justify-end">
                    <button
                        className="py-2.5 px-5 mr-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
                        onClick={() => setOpenModalForm("default")}>
                        Try again
                    </button>
                <button
                    className="py-2.5 px-5 mr-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
                    onClick={exportForm}>
                    Export
                </button>
            </div>
        </div>
    )
}

export default HeaderForm;