import { observer } from 'mobx-react-lite'
import { useState } from 'react'
import { useStoreContext } from '../../store/store'

interface PropType {
  product: {
    id: number
    name: string
    price: number
    units: number
    imageUrl: string
  }
}

const Product: React.FC<PropType> = observer(({ product }) => {
  const { cartStore } = useStoreContext()
  const [units, setUnits] = useState<number>(product.units)
  return (
    <div className="field">
      <div className="columns is-primary">
        <div className="column">
          <img src={product.imageUrl} alt="product placeholder"></img>
        </div>
        <div className="column">
          <p className="title is-3">{product.name}</p>
          <p className="title is-3">{product.price}</p>
        </div>
      </div>
      <input
        type="number"
        className="input"
        value={units}
        onChange={(e) => {
          let modifiedCount = parseInt(e.target.value)
          if (modifiedCount > 10) modifiedCount = 10
          if (modifiedCount < 1) modifiedCount = 1
          setUnits(modifiedCount)
          if (isNaN(modifiedCount)) return
          cartStore?.updateUnits(product.id, modifiedCount)
        }}
      />
    </div>
  )
})

export default Product
