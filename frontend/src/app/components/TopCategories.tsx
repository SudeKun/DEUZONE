import React from "react";

const TopCategories: React.FC = () => {
    return (
        <div className="flex flex-row w-11/12 h-screen gap-4 p-4 bg-gray-100">
            <div className="flex flex-col w-1/3 h-full md:w-1/2 gap-4">
                <div className="flex w-full h-2/3 rounded-md items-end justify-end border-2 pr-4 pb-4 border-borderprimary">
                    <h1 className="text-2xl font-semibold">Elektronik</h1>
                </div>
                <div className="flex flex-row w-full h-1/3 gap-4">
                    <div className="flex items-end justify-end pr-4 pb-4 w-1/2 h-full rounded-md border-2 border-borderprimary">
                        <h1 className="text-2xl font-semibold">Elektronik</h1>
                    </div>
                    <div className="flex items-end justify-end pr-4 pb-4 w-1/2 h-full rounded-md border-2 border-borderprimary">
                        <h1 className="text-2xl font-semibold">Elektronik</h1>
                    </div>
                </div>
            </div>

            <div className="flex flex-col w-2/3 h-full md:w-1/2 gap-4">
                <div className="flex items-end justify-end pr-4 pb-4 w-full h-1/3 rounded-md border-2 border-borderprimary">
                    <h1 className="text-2xl font-semibold">Elektronik</h1>
                </div>
                <div className="flex items-end justify-end pr-4 pb-4 w-full h-1/3 rounded-md border-2 border-borderprimary">
                    <h1 className="text-2xl font-semibold">Elektronik</h1>
                </div>
                <div className="flex items-end justify-end pr-4 pb-4 w-full h-1/3 rounded-md border-2 border-borderprimary">
                    <h1 className="text-2xl font-semibold">Elektronik</h1>
                </div>
            </div>
        </div>
    );
};

export default TopCategories;
