import type { VoidReturnAsyncFunc } from "./types"

export function onCtrlPlusS(func: VoidReturnAsyncFunc) {
    document.addEventListener(
        "keydown",
        function (e) {
            if (e.key === "s" && (e.metaKey || e.ctrlKey)) {
                e.preventDefault()
                func()
            }
        },
        false
    )
}
