import React from 'react'
import { Link } from 'react-router-dom'
interface Product {
  id: number
  name: string
  price: number
  imageUrl: string
}

interface PropType {
  product: Product
}

const ProductCard: React.FC<PropType> = ({ product }) => {
  const { id, name, price, imageUrl } = product

  return (
    <Link to={`/product?id=${id}`}>
      <div
        className="column is-full m-2 has-background-white"
        style={{ width: 250, height: 350 }}
      >
        <div className="card-image">
          <figure className="image is-1by1">
            <img width={200} height={200} src={imageUrl} alt="Placeholder" />
          </figure>
        </div>
        <div className="card-content">
          <p className="title is-4">{name}</p>
          <p className="subtitle is-5">₹{price}</p>
        </div>
      </div>
    </Link>
  )
}

export default ProductCard
