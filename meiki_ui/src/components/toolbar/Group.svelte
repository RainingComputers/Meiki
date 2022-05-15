<script lang="ts">
    export let show = true
</script>

<div class="flex flex-row gap-2" class:hidden={!show}>
    <slot />
</div>
