import { makeAutoObservable } from 'mobx'
import { AuthStore } from './AuthStore'
import React, { createContext, useContext } from 'react'
import { ProductStore } from './ProductStore'
import { CartStore } from './CartStore'

export class RootStore {
  authStore: AuthStore | null = null
  productStore: ProductStore | null = null
  cartStore: CartStore | null = null
  constructor() {
    makeAutoObservable(this)
    this.authStore = new AuthStore(this)
    this.productStore = new ProductStore(this)
    this.cartStore = new CartStore(this)
  }
}

export const StoreContext = createContext<RootStore>({} as RootStore)

export const useStoreContext = () => useContext(StoreContext)

export const RootStateProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  return (
    <StoreContext.Provider value={new RootStore()}>
      {children}
    </StoreContext.Provider>
  )
}
