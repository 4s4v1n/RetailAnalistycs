import { FC, useContext } from 'react';
import { Context } from '../..';

export interface HeadersProps {
    header?: string;
    mutationUpdate: any;
    mutationExport: any;
}

const HeaderView: FC<HeadersProps> = ({ header, mutationUpdate, mutationExport }) => {
    const { userStore } = useContext(Context)

    const acceptUpdate = () => {
        mutationUpdate()
    }

    return (
        <div className="flex justify-between sm:ml-64 p-4 pt-20">
            <h1 className="text-4xl font-extrabold dark:text-white">{header}</h1>
            <div className="flex justify-end">
                {userStore.role === "admin" ?
                    <button
                        className="py-2.5 px-5 mr-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
                        onClick={acceptUpdate}>
                        Update
                    </button>
                    : null}
                <button
                    className="py-2.5 px-5 mr-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
                    onClick={mutationExport}>
                    Export
                </button>
            </div>
        </div>
    )
}

export default HeaderView;