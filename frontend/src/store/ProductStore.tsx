import { makeAutoObservable } from 'mobx'
import { RootStore } from './store'

export interface Product {
  id: number
  name: string
  price: number
  imageUrl: string
  seller?: string
}

const DefaultProducts: Product[] = [
  {
    id: 1,
    name: 'T-shirt',
    price: 200,
    imageUrl: 'https://bulma.io/images/placeholders/128x128.png',
  },
  {
    id: 2,
    name: 'T-shirt',
    price: 300,
    imageUrl: 'https://bulma.io/images/placeholders/128x128.png',
  },
  {
    id: 3,
    name: 'T-shirt',
    price: 400,
    imageUrl: 'https://bulma.io/images/placeholders/128x128.png',
  },
  {
    id: 4,
    name: 'shirt',
    price: 200,
    imageUrl: 'https://bulma.io/images/placeholders/128x128.png',
  },
]

export class ProductStore {
  rootStore: RootStore | null = null
  products: Product[] | null = null
  constructor(rootStore: RootStore) {
    makeAutoObservable(this)
    this.rootStore = rootStore
    this.products = DefaultProducts
  }
  get(id: number): Product | undefined {
    const product = this.rootStore?.productStore?.products?.find(
      (product) => product.id == id
    )
    return product
  }
}
