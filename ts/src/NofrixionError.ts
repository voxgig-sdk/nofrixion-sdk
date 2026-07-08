
import { Context } from './Context'


class NofrixionError extends Error {

  isNofrixionError = true

  sdk = 'Nofrixion'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  NofrixionError
}

