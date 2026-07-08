
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { NofrixionSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await NofrixionSDK.test()
    equal(null !== testsdk, true)
  })

})
