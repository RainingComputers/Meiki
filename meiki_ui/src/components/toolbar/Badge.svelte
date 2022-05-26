<script context="module" lang="ts">
    export enum BadgeType {
        INFO,
        ERROR,
    }

    function getFontColorClass(type: BadgeType) {
        switch (type) {
            // TODO: Vishnu adds lipstick to this
            case BadgeType.INFO:
                return "text-gray-200 font-light"
            case BadgeType.ERROR:
                return "text-red-500 font-bold"
        }
    }

    function getIconColorClass(type: BadgeType) {
        switch (type) {
            // TODO: Vishnu adds lipstick to this
            case BadgeType.INFO:
                return "stroke-gray-200"
            case BadgeType.ERROR:
                return "stroke-red-500"
        }
    }
</script>

<script lang="ts">
    export let type: BadgeType = BadgeType.INFO
    export let label: string = ""
</script>

<span
    class="flex text-xs w-36 overflow-hidden px-3 gap-1 items-center {getFontColorClass(
        type
    )}"
>
    <span class="stroke-gray-200 h-3 w-3 {getIconColorClass(type)}">
        <slot />
    </span>
    {label}
</span>
