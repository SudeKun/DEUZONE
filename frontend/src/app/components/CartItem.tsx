'use client'
import React, { useState } from 'react'
import { FaPlus, FaMinus } from "react-icons/fa6"
import { FaRegTrashAlt } from "react-icons/fa"

export interface CartItemProps {
    Item: {
        name: string,
        price: number,
        rating: number,
        quantity: number,
        image: string,
    }
}

const CartItem: React.FC<CartItemProps> = ({ Item }) => {

    const [quantity, setQuantity] = useState(Item.quantity);

    const plusButton = () => {
        setQuantity(quantity + 1);
    }

    const minusButton = () => {
        if (quantity > 1)
            setQuantity(quantity - 1);
    }

    return (
        <div className='w-4/5 h-40 flex flex-row border-2 rounded-lg'>
            <div className='h-full w-1/5 aspect-[1/1] flex justify-center items-center'>
                <img className='object-cover w-full h-full' src={Item.image} alt={Item.name} />
            </div>
            <div className="h-full w-3/5 px-4 flex flex-col justify-around">
                <h1 className='font-semibold text-xl'>{Item.name}</h1>
                <div className="flex flex-row w-32 h-10 border-2 rounded-full justify-around items-center">
                    <FaPlus onClick={plusButton} className='cursor-pointer' />
                    {quantity}
                    <FaMinus onClick={minusButton} className='cursor-pointer'/>
                </div>
                <h1>${quantity * Item.price}</h1>
            </div>
            <div className='w-1/5 h-full flex justify-center items-center text-red hover:bg-red hover:text-white cursor-pointer transition-all duration-100'>
                <FaRegTrashAlt />
            </div>
        </div>
    )
}

export default CartItem;
