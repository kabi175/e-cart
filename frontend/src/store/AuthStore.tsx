import { runInAction, makeAutoObservable } from 'mobx'
import { RootStore } from './store'
import { login, signup } from '../api/user'
export class AuthStore {
  userName: string | null = null
  rootStore: RootStore | null = null

  constructor(rootStore: RootStore) {
    makeAutoObservable(this)
    this.rootStore = rootStore
    this.userName = null
  }

  async signup(email: string, userName: string, password: string) {
    const username = await signup(email, password, userName)

    runInAction(() => {
      if (this.rootStore?.authStore != null) {
        this.rootStore.authStore.userName = username
      }
    })
  }
  async login(email: string, password: string) {
    const username = await login(email, password)
    console.log(username)
    runInAction(() => {
      if (username != null && this.rootStore?.authStore != null) {
        this.rootStore.authStore.userName = username
      }
    })
  }

  signOut(): void {
    // sign out
  }
}
