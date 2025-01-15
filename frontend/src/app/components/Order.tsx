'use client'
import React from 'react'
import { useState } from 'react'

interface Order {
    Data: {
        order_id: string,
        customer_id: string,
        cart_id: string,
        status: boolean,
        date: string,
        price: number
    }
}

const products = [
    { id: 1, price: 100, quantity: 2 },
    { id: 2, price: 200, quantity: 3 }
]

const Order: React.FC<Order> = ({ Data }) => {

    const [status, setStatus] = useState(false);

    const handleClick = () => {
        setStatus(!status);
    }

    return (
        <div onClick={handleClick} className='h-40 mt-4 w-11/12 border-2 rounded-md shadow-lg cursor-pointer'>
            <div className="h-full w-full flex flex-row">
                <div className="h-full w-1/2 flex flex-col justify-center pl-2">
                    <h1 className='font-bold text-md'>Order ID: {Data.order_id}</h1>
                    <h1 className='font-bold text-md'>Date: {Data.date}</h1>
                </div>
                <div className="h-full w-1/2 flex flex-col justify-center items-end pr-8">
                    <h1 className='font-bold text-md'>${Data.price}</h1>
                </div>
            </div>

            {
                status && (
                    <div className="flex flex-col mt-2 w-full border-2 rounded-md shadow-inner">
                        <ul>
                            {
                                products.map((p) => (
                                    <li key={p.id} className='w-full h-40 flex flex-row border-t-2'>
                                        <div className="h-full w-1/2 flex items-centerr pl-2">
                                            <h1 className='font-bold text-md'>Order ID: {p.id}</h1>
                                        </div>
                                        <div className="h-full w-1/2 flex flex-col justify-center pr-2">
                                            <h1 className='font-bold text-md'>${p.price}</h1>
                                            <button className='bg-primary text-textColor rounded-md w-1/4'>Create a Review</button>
                                        </div>
                                    </li>
                                ))
                            }
                        </ul>
                    </div>
                )
            }
        </div>
    )
}

export default Order
