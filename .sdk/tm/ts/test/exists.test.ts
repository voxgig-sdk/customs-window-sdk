
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { CustomsWindowSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await CustomsWindowSDK.test()
    equal(null !== testsdk, true)
  })

})
