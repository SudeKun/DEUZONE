import Order from '@/app/components/Order'
import React from 'react'

const order: Order = {
    Data: {
        order_id: "12345",
        customer_id: "67890",
        cart_id: "abc123",
        status: true,
        date: "2025-01-14T12:00:00Z",
        price: 300,
    }
};

const OrderHistory : React.FC = () => {
    return (
        <div className="container h-screen flex flex-col items-center">
            <h1 className='font-bold text-2xl mt-4'>Order History</h1>
            <Order Data={order.Data} />
            <Order Data={order.Data} />
        </div>
    )
}

export default OrderHistory;
