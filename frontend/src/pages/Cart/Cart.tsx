import { observer } from 'mobx-react-lite'
import React from 'react'
import { Link } from 'react-router-dom'
import { useStoreContext } from '../../store/store'
import Product from './Product'

const OrderDetails: React.FC = observer(() => {
  const { cartStore } = useStoreContext()
  const totalAmount = cartStore?.cart?.reduce(
    (totalAmount, { price, units }) => {
      totalAmount += price * units
      return totalAmount
    },
    0
  )

  return (
    <div className="column mt-6">
      {' '}
      <div className="box">
        <p className="title is-4 is-uppercase  has-text-weight-semibold">
          Price Details
        </p>
        <p className="title is-5 has-text-weight-semibold">
          Total Amount : {totalAmount}{' '}
        </p>
      </div>
    </div>
  )
})

const Cart: React.FC = observer(() => {
  const { cartStore } = useStoreContext()
  return (
    <div className="has-background-light">
      <div className="container">
        <div className="columns is-variable is-8-mobile is-8-tablet is-8-desktop is-8-widescreen is-8-fullhd is-multiline is-centered is-mobile ">
          <div className="column mt-6">
            <div className="columns box is-multiline is-centered ">
              <p className="column is-uppercase  has-text-weight-bold is-full title">
                My Cart
              </p>
              <div>
                <div className="column is-full">
                  {cartStore?.cart?.map((product, id) => (
                    <Product key={id} product={product} />
                  ))}
                </div>
                <div className="column is-full">
                  <div className="columns is-centered">
                    <div className="column is-narrow">
                      <button className="button is-info">PLACE ORDER</button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          {<OrderDetails />}
        </div>
      </div>
    </div>
  )
})

const EmptyCary: React.FC = () => {
  return (
    <div className="column is-full ">
      <p className="title">Your Cart is Empty</p>
      <Link to="/" className="button">
        Shop Now
      </Link>
    </div>
  )
}

const CartContanier: React.FC = () => {
  const { cartStore } = useStoreContext()
  if (
    cartStore == null ||
    cartStore.cart == null ||
    cartStore.cart.length == 0
  ) {
    return <EmptyCary />
  }
  return <Cart />
}

export default observer(CartContanier)
