import { makeAutoObservable } from 'mobx'
import { RootStore } from './store'

export class AuthStore {
  userName: string | null = null
  rootStore: RootStore | null = null

  constructor(rootStore: RootStore) {
    makeAutoObservable(this)
    this.userName = 'kabi'
    this.rootStore = rootStore
  }

  signIn(userName: string, password: string): void {
    //sign in
    if (this.rootStore?.authStore) {
      this.rootStore.authStore.userName = userName
    }
  }

  signOut(): void {
    // sign out
  }
}
