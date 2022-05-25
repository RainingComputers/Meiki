// Based on https://github.com/component/debounce/blob/master/index.js

import type { VoidReturnAsyncFunc } from "./types"

const DEFAULT_WAIT_TIME_MS = 500

export function debounce(
    func: VoidReturnAsyncFunc,
    wait: number = DEFAULT_WAIT_TIME_MS
): VoidReturnAsyncFunc {
    let timeout: NodeJS.Timeout
    let args: IArguments
    let context: Function
    let timestamp: number
    let result: Promise<void>

    function later() {
        const last = Date.now() - timestamp

        if (last < wait && last >= 0) {
            timeout = setTimeout(later, wait - last)
        } else {
            timeout = null
            result = func.apply(context, args)
            context = args = null
        }
    }

    async function debounced() {
        context = this
        args = arguments
        timestamp = Date.now()

        if (!timeout) {
            timeout = setTimeout(later, wait)
            result = await func.apply(context, args)
            context = args = null
        }

        return result
    }

    return debounced
}
