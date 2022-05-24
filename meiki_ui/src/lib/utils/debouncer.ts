type VoidReturnAsyncFunc = (...args: any[]) => Promise<void>

// Based on https://github.com/component/debounce/blob/master/index.js

export function debounce(
    func: VoidReturnAsyncFunc,
    wait: number = 500,
    immediate: boolean = false
): VoidReturnAsyncFunc {
    var timeout: NodeJS.Timeout,
        args: any,
        context: any,
        timestamp: number,
        result: Promise<void>
    if (null == wait) wait = 100

    function later() {
        var last = Date.now() - timestamp

        if (last < wait && last >= 0) {
            timeout = setTimeout(later, wait - last)
        } else {
            timeout = null
            if (!immediate) {
                result = func.apply(context, args)
                context = args = null
            }
        }
    }

    var debounced = async function () {
        context = this
        args = arguments
        timestamp = Date.now()
        var callNow = immediate && !timeout
        if (!timeout) timeout = setTimeout(later, wait)
        if (callNow) {
            result = await func.apply(context, args)
            context = args = null
        }

        return result
    }

    return debounced
}
