package utiltools

//具体的API可以参见GO-CHART组件
//https://github.com/wcharczuk/go-chart
import (
	"bytes"
	"commonlib/go-chart"
	"io/ioutil"
)

//图片宽度高度配置
const (
	chartWidth  int = 580
	chartHeight int = 290
)

//二给拆线图
/*
xTicks []chart.Tick X轴刻度
yTicks []chart.Tick Y轴刻度
minYFloat float64
maxYFloat float64
series []chart.Series 拆线数据点集合
chartName string
dir string 图片存储目录
xName X轴名称
yName Y轴名称
*/
func DrawCurveChart(xTicks []chart.Tick, yTicks []chart.Tick,
	minYFloat float64, maxYFloat float64,
	series []chart.Series, chartName string,
	dir string, xName string, yName string) error {
	graph := chart.Chart{
		Width:  chartWidth,
		Height: chartHeight,
		XAxis: chart.XAxis{
			Name:      xName,
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			Ticks:     xTicks,
			//显示相应的网格，X轴对应的是Y坐标
			GridMajorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorBlackTransparent,
			},
			GridMinorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorBlackTransparent,
			},
		},
		YAxis: chart.YAxis{
			Name:      yName,
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			Range:     &chart.ContinuousRange{Min: minYFloat, Max: maxYFloat},
			Ticks:     yTicks,
			//显示相应的网格，Y轴对应的是X坐标
			GridMajorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorBlackTransparent,
			},
			GridMinorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorBlackTransparent,
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:    10,
				Left:   20,
				Bottom: 10,
				Right:  20,
			},
			FillColor: chart.ColorBlackTransparentBackground,
		},
		Series: series,
	}
	//添加LEGEND图例标识
	graph.Elements = []chart.Renderable{
		chart.LegendThin(&graph),
	}
	//写文件
	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		return err
	}
	path := dir + chartName
	err = ioutil.WriteFile(path, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
