
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { CustomsWindowSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = CustomsWindowSDK.test()
    equal(testsdk instanceof CustomsWindowSDK, true,
      'CustomsWindowSDK.test() must return a client synchronously')
  })

})
