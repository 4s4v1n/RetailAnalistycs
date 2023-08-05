import { FC, useContext, useRef, useState } from 'react';
import { Context } from '../..';
import AddModal from '../Modal/AddModal';
import Papa from 'papaparse'

export interface HeadersProps {
    header?: string;
    keys: string[];
    mutationAdd: any;
    mutationExport: any;
    mutationImport: any;
}

const Header: FC<HeadersProps> = ({ header, keys, mutationAdd, mutationExport, mutationImport }) => {
    const { userStore } = useContext(Context)
    const [openModalAdd, setOpenModalAdd] = useState<string | undefined>();

    const inputRef = useRef<HTMLInputElement>(null)

    const selectFile = (event: React.ChangeEvent<HTMLInputElement>) => {
        const files =  event.currentTarget.files;
        if (files === null) {
            return
        }
        const file = files.item(0)
        if (file === null) {
            return
        }
        Papa.parse(file, {
            header: true,
            download: true,
            skipEmptyLines: true,
            delimiter: "\t",
            complete: (results) => {
                mutationImport(results.data)
            },
          })
    };

    const handleClick = () => {
        inputRef.current?.click()
    };

    return (
        <div className="flex justify-between sm:ml-64 p-4 pt-20">
            <h1 className="text-4xl font-extrabold dark:text-white">{header}</h1>
            <AddModal openModalAdd={openModalAdd} setOpenModalAdd={setOpenModalAdd} keys={keys} mutationAdd={mutationAdd} />
            <div className="flex justify-end">
                {userStore.role === "admin" ?
                    <div>
                        <button
                            className="py-2.5 px-5 mr-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
                            onClick={() => setOpenModalAdd("default")}>
                            Add
                        </button>
                        <>
                            <input id="fileImport" ref={inputRef} type="file" className="hidden" accept=".tsv" onChange={selectFile} />
                            <button
                                className="py-2.5 px-5 mr-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
                                onClick={handleClick}
                            >
                                Import
                            </button>
                        </>
                    </div>
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

export default Header;