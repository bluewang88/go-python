'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _reactDnd = require('react-dnd');

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var DragTypes = {
  HEADER: 'header'
};
var headerSource = {
  beginDrag: function beginDrag(_ref) {
    var label = _ref.label;

    return { label: label };
  }
};
var headerTarget = {
  hover: function hover(targetProps, monitor) {
    var targetLabel = targetProps.label;
    var sourceProps = monitor.getItem();
    var sourceLabel = sourceProps.label;

    if (sourceLabel !== targetLabel && targetProps.onMove) {
      targetProps.onMove({ sourceLabel: sourceLabel, targetLabel: targetLabel });
    }
  },
  drop: function drop(targetProps) {
    if (targetProps.onFinishMove) {
      targetProps.onFinishMove();
    }
  }
};

var dragSource = (0, _reactDnd.DragSource)( // eslint-disable-line new-cap
DragTypes.HEADER, headerSource, function (connect) {
  return {
    connectDragSource: connect.dragSource()
  };
});
var dropTarget = (0, _reactDnd.DropTarget)( // eslint-disable-line new-cap
DragTypes.HEADER, headerTarget, function (connect) {
  return {
    connectDropTarget: connect.dropTarget()
  };
});
var header = function header(_ref2) {
  var connectDragSource = _ref2.connectDragSource,
      connectDropTarget = _ref2.connectDropTarget,
      label = _ref2.label,
      children = _ref2.children,
      onMove = _ref2.onMove,
      onFinishMove = _ref2.onFinishMove,
      props = _objectWithoutProperties(_ref2, ['connectDragSource', 'connectDropTarget', 'label', 'children', 'onMove', 'onFinishMove']);

  return connectDragSource(connectDropTarget(_react2.default.createElement(
    'th',
    props,
    children
  )));
};

exports.default = dragSource(dropTarget(header));