markdownfmt = require('./markdownfmt.js')

class MarkdownFormat
  constructor: ->
    atom.workspace.observeTextEditors (editor) =>
      @handleSave(editor)

  destroy: ->
    @subscription.dispose()

  handleSave: (editor) ->
    buffer = editor.getBuffer()
    @subscription = buffer.onWillSave =>
      if editor.getGrammar().scopeName is 'source.gfm'
        buffer.setTextViaDiff(markdownfmt.ProcessMarkdown(buffer.getText()))

module.exports =
  activate: ->
    @markdownFormat = new MarkdownFormat()

  deactivate: ->
    @markdownFormat.destroy()
