import { observer } from 'mobx-react-lite'
import React from 'react'
import { useStoreContext } from '../../store/store'
import ProductCard from './ProductCard'

const Home: React.FC = () => {
  const { productStore } = useStoreContext()
  return (
    <div className="has-background-light">
      <div className="pt-6 container">
        <div className="columns is-mobile is-multiline is-centered">
          {productStore?.products?.map((product, id) => (
            <ProductCard key={id} product={product} />
          ))}
        </div>
      </div>
    </div>
  )
}

export default observer(Home)
