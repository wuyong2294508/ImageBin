package main

import (
	"fmt"
	"os"
	"image"
	"image/jpeg"
	_ "image/png"
	"image/color"
	"math"
)

func main() {
	leng := len(os.Args);
	if leng <= 1 {
		fmt.Println("请重新运行程序，并在命令行后输入图片的完整路径！");
		return;
	}
	fmt.Println(os.Args[1]);
	reader, err := os.Open(os.Args[1]);
	if err != nil {
		return;
	}
	defer reader.Close();

	img, _, err := image.Decode(reader);

	//获取图像的长宽
	bounds := img.Bounds();
    dx := bounds.Dx();
    dy := bounds.Dy();
	
    newGray := image.NewGray(bounds);
	var temp [][]uint8;
	//转化为灰度图像,并存储到二维切片temp中。
    for i := 0; i < dx; i++ {
		s := make([]uint8,0,dy)
        for j := 0; j < dy; j++ {
            colorRgb := img.At(i, j);
            r, g, b, _ := colorRgb.RGBA();
			//没有弄明白这里要右移8位的原因
			rgb := uint8((r>>8 + g>>8 + b>>8)/3);
			s = append(s,rgb);
        }
		temp = append(temp,s);
    }
	
	fmt.Println(dx);
	fmt.Println(dy);
	
	dx_f64 := float64(dx);
	dy_f64 := float64(dy);
	height := math.Sqrt(dy_f64/dx_f64*(dy_f64+dx_f64));
	width := dx_f64/dy_f64*height;
	
	fmt.Println(height);
	fmt.Println(width);
	
	for h := 0; h < dx; h++ {
        for w := 0; w < dy; w++ {
			h1 := h - int(math.Floor(height/2));
			h2 := h + int(math.Floor(height/2));
			w1 := w - int(math.Floor(width/2));
			w2 := w + int(math.Floor(width/2));
			
			if h1 < 0 {
				h1 = 0;
			}
			if h2 > dx {
				h2 = dx;
			}
			if w1 < 0 {
				w1 = 0;
			}
			if w2 > dy {
				w2 = dy;
			}
			
			//fmt.Println(h1);
			//fmt.Println(h2);
			//fmt.Println(w1);
			//fmt.Println(w2);
			
			sum := 0;
			num := 0;
			for i := h1;i < h2;i++ {
				for j := w1;j < w2;j++ {
					sum += int(temp[i][j]);
					num += 1;
				}
			}
			mean := sum/num;
			grey := temp[h][w];
			if(int(grey) < int(float64(mean)*0.9)) {
				grey = 0;;
			}else{
				grey = 255;
			}
			
			newGray.SetGray(h, w, color.Gray{grey});
		}
	}
	
	outpath := "result.jpg";
	outfile, err := os.Create(outpath);
	if err != nil {
		return;
	}
	defer outfile.Close();

	jpeg.Encode(outfile, newGray, nil);

}















