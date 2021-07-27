import { makeAutoObservable } from 'mobx'
import { Product } from './ProductStore'
import { RootStore } from './store'

interface CartProduct extends Product {
  units: number
}

export class CartStore {
  rootStore: RootStore | null = null
  cart: CartProduct[] | null = null
  constructor(rootStore: RootStore) {
    makeAutoObservable(this)
    this.rootStore = rootStore
  }

  add(product: Product) {
    if (this.rootStore?.cartStore && this.rootStore.cartStore.cart == null) {
      this.rootStore.cartStore.cart = [{ ...product, units: 1 }]
      return
    }
    const result = this.rootStore?.cartStore?.cart?.find(
      (p) => p.id == product.id
    )
    if (result != undefined) return
    this.rootStore?.cartStore?.cart?.push({ ...product, units: 1 })
  }

  remove(id: number) {
    const newCart = this.rootStore?.cartStore?.cart?.filter(
      (product) => product.id != id
    ) as CartProduct[]
    if (this.rootStore?.cartStore?.cart) this.rootStore.cartStore.cart = newCart
  }
  updateUnits(id: number, newUnits: number) {
    const newCart = this.rootStore?.cartStore?.cart?.map((product) => {
      if (product.id == id) product.units = newUnits
      return product
    }) as CartProduct[]
    if (this.rootStore?.cartStore?.cart) this.rootStore.cartStore.cart = newCart
  }
}
