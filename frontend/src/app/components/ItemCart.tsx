import React from 'react'

interface ItemCartProps {
  title: string;
  price: number;
  image: string;
  stars: number;
  comments: number;
}

const ItemCart: React.FC<ItemCartProps> = ({title, price, image, stars, comments}) => {
  return (
    <div className='flex flex-col items-center relative group w-52 h-80 rounded-2xl border-black border-2 overflow-hidden'>
      <div className='flex items-center justify-center h-[60%] w-[100%]'>
        <img src={image} alt={title} className='w-full h-full object-cover' />
      </div>
      <div className='flex flex-col justify-between items-start pl-4 py-3 h-[40%] w-[100%] mb-[36px]'>
        <h1>{title}</h1>
        <h3>{stars} ({comments})</h3>
        <h2>{price}</h2>
      </div>
      <button className='bg-primary absolute bottom-1 w-[90%] h-[36px] mb-[2px] rounded-full hidden group-hover:block'>Add to Cart</button>
    </div>
  )
}

export default ItemCart;
