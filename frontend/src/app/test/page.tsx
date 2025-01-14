import React from 'react'
import CartItem, { CartItemProps } from '../components/CartItem'
import { Product } from '../lib/items/ProductList'

const produc: CartItemProps = {
    Item: {
        name: "Wodaaa",
        price: 5000,
        rating: 4.5,
        quantity: 4,
        image: "X"
    }
}


const page = () => {
  return (
    <div>
      <CartItem Item={produc.Item}/>
    </div>
  )
}

export default page
