import React from 'react'
import ProductCard from './ProductCard'
const Home: React.FC = () => {
  return (
    <div className="m-6 container has-background-lite">
      <div className="columns is-mobile is-multiline is-centered">
        <ProductCard product={{ id: 1, name: 'T-Shirt', price: 200 }} />
        <ProductCard product={{ id: 2, name: 'T-Shirt', price: 200 }} />
        <ProductCard product={{ id: 3, name: 'T-Shirt', price: 200 }} />
        <ProductCard product={{ id: 4, name: 'T-Shirt', price: 200 }} />
        <ProductCard product={{ id: 1, name: 'T-Shirt', price: 200 }} />
        <ProductCard product={{ id: 1, name: 'T-Shirt', price: 200 }} />
        <ProductCard product={{ id: 1, name: 'T-Shirt', price: 200 }} />
        <ProductCard product={{ id: 1, name: 'T-Shirt', price: 200 }} />
        <ProductCard product={{ id: 1, name: 'T-Shirt', price: 200 }} />
        <ProductCard product={{ id: 1, name: 'T-Shirt', price: 200 }} />
      </div>
    </div>
  )
}

export default Home
