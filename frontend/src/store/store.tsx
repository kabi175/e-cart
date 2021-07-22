import { makeAutoObservable } from 'mobx'
import { AuthStore } from './AuthStore'
import React, { createContext } from 'react'

export class RootStore {
  authStore: AuthStore | null = null

  constructor() {
    makeAutoObservable(this)
    this.authStore = new AuthStore(this)
  }
}

export const StoreContext = createContext<RootStore>({} as RootStore)

export const RootStateProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  return (
    <StoreContext.Provider value={new RootStore()}>
      {children}
    </StoreContext.Provider>
  )
}
