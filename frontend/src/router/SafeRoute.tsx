import React, { useContext } from 'react'
import { observer } from 'mobx-react-lite'
import { Redirect } from 'react-router-dom'
import { StoreContext } from '../store/store'

interface PropType {
  children: React.ReactNode
}

const SafeRoute: React.FC<PropType> = ({ children }) => {
  const { authStore } = useContext(StoreContext)
  if (authStore == null || authStore.userName == null) {
    return <Redirect to="/login" />
  }
  return <>{children}</>
}

export default observer(SafeRoute)
