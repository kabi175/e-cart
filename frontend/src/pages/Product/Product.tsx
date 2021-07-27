import React from 'react'
import { useLocation } from 'react-router-dom'
import { useStoreContext } from '../../store/store'
import { observer } from 'mobx-react-lite'
const Product: React.FC = () => {
  const { productStore, cartStore } = useStoreContext()

  function useQuery() {
    return new URLSearchParams(useLocation().search)
  }

  const productId = useQuery().get('id')
  if (!productId) {
    return <></>
  }

  const product = productStore?.get(parseInt(productId))
  if (!product) {
    return <></>
  }

  const { name, price, imageUrl } = product

  return (
    <div className="has-background-light p-6">
      <div className="container ">
        <div className="columns is-multiline is-mobile">
          <div className="column" style={{ maxWidth: 500 }}>
            <figure className="image is-fullwidth">
              <img src={imageUrl} alt="Product_Image" />
            </figure>
          </div>
          <div className="column " style={{ minWidth: 300 }}>
            <p className="title"> {name} </p>
            <p className="title"> ₹{price} </p>
            <p className="buttons">
              <button
                className="button is-link is-outlined"
                onClick={() => {
                  cartStore?.add(product)
                }}
              >
                ADD TO CART
              </button>
              <button className="button is-link">BUY NOW</button>
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default observer(Product)
