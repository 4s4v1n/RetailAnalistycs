import { FC } from "react";
import { RxCross2, RxInfoCircled } from "react-icons/rx";


export interface ModalDeleteProps {
    openModalDelete?: string | undefined
    setOpenModalDelete: React.Dispatch<React.SetStateAction<string | undefined>>
    mutationDelete: any;
    selectedValue: any;
}

const DeleteModal: FC<ModalDeleteProps> = ({ openModalDelete, setOpenModalDelete, mutationDelete, selectedValue }) => {

    const acceptDelete = () => {
        mutationDelete(selectedValue)
    }

    return (
        <div>
            {
                openModalDelete !== undefined ? (
                    <div className="flex justify-center items-center overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none">
                        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>
                        <div className="relative w-full max-w-md max-h-full">
                            <div className="relative bg-white rounded-lg shadow dark:bg-gray-700">
                                <button onClick={() => setOpenModalDelete(undefined)}
                                    className="absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-800 dark:hover:text-white"
                                    data-modal-hide="popup-modal">
                                    <RxCross2 className="w-5 h-5 fill-white" />
                                </button>
                                <div className="p-6 text-center">
                                    <RxInfoCircled className="mx-auto mb-4 text-gray-400 w-14 h-14 dark:text-gray-200" />
                                    <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
                                        Are you sure you want to delete this item?
                                    </h3>
                                    <button onClick={() => { acceptDelete(); setOpenModalDelete(undefined) }}
                                        data-modal-hide="popup-modal"
                                        className="text-white bg-red-600 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 dark:focus:ring-red-800 font-medium rounded-lg text-sm inline-flex items-center px-5 py-2.5 text-center mr-2">
                                        Yes, I'm sure
                                    </button>
                                    <button onClick={() => setOpenModalDelete(undefined)}
                                        data-modal-hide="popup-modal"
                                        className="text-gray-500 bg-white hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-gray-200 rounded-lg border border-gray-200 text-sm font-medium px-5 py-2.5 hover:text-gray-900 focus:z-10 dark:bg-gray-700 dark:text-gray-300 dark:border-gray-500 dark:hover:text-white dark:hover:bg-gray-600 dark:focus:ring-gray-600">
                                        No, cancel
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                ) : null
            }
        </div>
    )
}

export default DeleteModal;