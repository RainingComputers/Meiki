<script lang="ts">
    // TODO style this if needed

    import Markdoc from "@markdoc/markdoc"

    export let text: string = ""

    $: ast = Markdoc.parse(text)
    $: content = Markdoc.transform(ast)
</script>

<p data-cy="renderer" class="py-3 pb-24 px-5 overflow-scroll h-full prose max-w-full">
    {@html Markdoc.renderers.html(content)}
</p>
